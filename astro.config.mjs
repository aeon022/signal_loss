import { defineConfig } from 'astro/config';
import tailwindcss from '@tailwindcss/postcss';

// https://astro.build/config
// GitHub Pages serves project sites (not a custom domain) from a
// /<repo>/ subpath — base must match, and every asset reference in the
// codebase uses import.meta.env.BASE_URL rather than a root-relative
// path, specifically so this works whether it's deployed here or later
// moved to a host that serves from the true root.
export default defineConfig({
  site: 'https://aeon022.github.io',
  base: '/signal_loss/',
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
