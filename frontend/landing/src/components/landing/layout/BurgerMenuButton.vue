<template>
  <button
    :class="{
      'burger-menu-button': true,
      active: menuState,
    }"
    aria-label="Navigation menu button"
    @click.prevent="toggleMenuState"
  >
    <span></span>
  </button>
</template>

<script lang="ts" setup>
import { useLandingMenuState } from '@/services/landing/menuState.js';

const { menuState, toggleMenuState } = useLandingMenuState();
</script>

<style lang="postcss">
.burger-menu-button {
  width: 38px;
  height: 38px;
  display: flex;
  justify-content: center;
  align-items: center;
  border: none;
  user-select: none;

  @apply min-lg:hidden;

  span {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100%;
    pointer-events: none;

    &::before,
    &::after {
      content: '';
      display: block;
      height: 1.4px;
      width: 22px;
      border-radius: 2px;
      background-color: theme('colors.gray.70');
      transition: transform 0.15s ease;
    }

    &::before {
      transform: translateY(-4px) rotate(0deg);
    }

    &::after {
      transform: translateY(4px) rotate(0deg);
    }
  }

  &.active span {
    &::before {
      transform: translateY(1px) rotate(45deg);
    }

    &::after {
      transform: translateY(0) rotate(-45deg);
    }
  }
}
</style>
