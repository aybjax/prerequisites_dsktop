<div>
  <div use:create></div>
</div>
 
<script>
  import SvelteRouter, {  } from 'svelte-router'
  import HomePage from './pages/ConnectionPage.svelte';
  import GraphPage from './pages/GraphPage.svelte';
 
  function create (node) {
    const router = new SvelteRouter({
      target: node,
      mode: 'hash',
      routes: [
        {
          path: '/graph',
          component: GraphPage,
        },
        {
          path: '/',
          component: HomePage
        },
    ]
    })
 
    return {
      destroy () {
        router.destroy()
      }
    }
  }
</script> 
 
<style>
</style>