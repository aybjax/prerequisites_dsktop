<div>
  <Link to="/">Home</Link>
  <Link to="/welcome">Welcome</Link>
  <div use:create></div>
</div>
 
<script>
  import SvelteRouter, { Link } from 'svelte-router'
  import Home from './pages/HomePage.svelte'
  import Welcome from './pages/AboutPage.svelte'
 
  function create (node) {
    const router = new SvelteRouter({
      target: node,
      mode: 'hash',
      routes: [{
        path: '/',
        component: Home
      }, {
        path: '/welcome',
        component: Welcome
      }]
    })
 
    return {
      destroy () {
        router.destroy()
      }
    }
  }
</script> 
 
<style>
  .router-link-active {
    color: red;
  }
</style> 


<!-- <script>
  import { Router } from 'svelte-router-spa'
  import { routes } from './routes'
</script> -->

<!-- <main>
  <Router {routes} />
  Hello
</main> -->