import { defineConfig } from 'astro/config';
import tailwindcss from '@tailwindcss/postcss';

// https://astro.build/config
export default defineConfig({
  output: 'static',
  devToolbar: {
    enabled: false,
  },
  vite: {
    css: {
      postcss: {
        plugins: [tailwindcss()],
      },
    },
  },
});
