import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  base: '/overlays',
  server: {
    host: true,
    port: 3008,
  },
	clearScreen: false,
});
