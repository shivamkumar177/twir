<template>
  <svg
    :viewBox="viewBox"
    :aria-label="label"
    :width="width"
    :height="height"
    :style="styles"
    :class="filteredClasses"
    :stroke="isOutline ? stroke : 'none'"
    :fill="isOutline ? 'none' : fill"
  >
    <title v-if="title">{{ title }}</title>
    <g>
      <path
        v-for="(path, index) in icon.path"
        :key="index"
        :d="path.d"
        :vector-effect="isOutline ? 'non-scaling-stroke' : undefined"
      ></path>
    </g>
  </svg>
</template>

<script lang="ts" setup>
import { computed, type CSSProperties } from 'vue';

import * as icons from './icons';
import type { IconName, Icon } from './types.js';

const props = withDefaults(
  defineProps<{
    name: IconName;
    width?: number;
    height?: number;
    title?: string;
    label?: string;
    strokeWidth?: number;
    stroke?: string;
    fill?: string;
    strokeLinejoin?: 'bevel' | 'round' | 'miter';
    strokeLinecap?: 'butt' | 'round' | 'square';
    class?: string;
    rotate?: number;
  }>(),
  {
    class: undefined,
    rotate: undefined,
    title: undefined,
    label: undefined,
    stroke: undefined,
    fill: undefined,
    width: 24,
    height: 24,
    strokeWidth: 1.5,
    strokeLinecap: 'round',
    strokeLinejoin: 'round',
  },
);

const icon = computed(() => icons[props.name]);
if (!icon.value) {
  throw new Error(`Cannot find icon "${props.name}"`);
}

const excludeStyleMap: Record<Icon['style'], string> = {
  'outline': 'fill',
  'solid': 'stroke',
};

const filteredClasses = computed(() => {
  if (!props.class) return undefined;

  const property = excludeStyleMap[icon.value.style];
  return props.class.split(' ').filter(cl => cl.includes(property) ? false : true).join(' ');
});

const viewBox = computed(() => `0 0 ${icon.value.width} ${icon.value.height}`);

const transform = computed(() => props.rotate ? `rotate(${props.rotate}deg)` : undefined);

const isOutline = computed(() => icon.value.style === 'outline');

const styles = computed<CSSProperties>(() => {
  return isOutline.value
    ? {
        strokeWidth: props.strokeWidth,
        strokeLinecap: props.strokeLinecap,
        strokeLinejoin: props.strokeLinejoin,
        transform: transform.value,
      }
    : { transform: transform.value };
});
</script>
