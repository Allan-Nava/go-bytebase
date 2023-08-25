import { defineConfig } from 'vitepress';
//
// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: 'go-bytebase',
  description: '',
  base: '/go-bytebase/',
  themeConfig: {
    outline: [2, 3],
    socialLinks: [
      { icon: 'github', link: 'https://github.com/Allan-Nava/go-bytebase' },
    ],
    sidebar: [
      { text: 'Introduction', link: '/' },
    ],
  },
})