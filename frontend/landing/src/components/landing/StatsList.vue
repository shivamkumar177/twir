<template>
  <div v-if="isLoading || !stats">
    <div class="w-full h-[70px]"></div>
  </div>
  <Flicking
    v-else-if="isRenderSlider"
    :plugins="plugins"
    :options="{
      panelsPerView: slidesPerView,
      align: 'next',
      bound: true,
    }"
    class="flex w-full max-sm:mx-0 animate-fadeIn opacity-0 cursor-grab select-none"
  >
    <StatsItem v-for="item in stats" :key="item.name" :item="item" class="w-full" />
  </Flicking>
  <div
    v-else
    class="
      inline-flex
      gap-x-6
      justify-between
      w-full
      max-md:opacity-0
      overflow-hidden
      animate-fadeIn
    "
  >
    <StatsItem v-for="item in stats" :key="item.name" :item="item" class="w-full flex-1" />
  </div>
</template>

<script lang="ts" setup>
import { AutoPlay } from '@egjs/flicking-plugins';
import Flicking from '@egjs/vue3-flicking';
import { useWindowSize } from '@vueuse/core';
import { computed } from 'vue';

import StatsItem from '@/components/landing/StatsItem.vue';
import { useStats } from '@/services/landing';

const { width: windowWidth } = useWindowSize();
const { data: stats, isLoading } = useStats();

const isRenderSlider = computed(() => windowWidth.value <= 768);

const slidesPerView = computed(() => {
  if (windowWidth.value < 410) {
    return 1;
  } else if (windowWidth.value < 568) {
    return 2;
  } else {
    return 3;
  }
});

const plugins = [new AutoPlay()];
</script>

<style>
@import '@egjs/vue3-flicking/dist/flicking.css';
</style>
