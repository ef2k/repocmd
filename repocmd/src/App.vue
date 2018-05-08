<template>
  <div id="app">
    <RepoList/>
  </div>
</template>

<script>
import RepoList from './components/RepoList'

export default {
  name: 'App',
  components: {
    RepoList
  }
}
</script>

<style>
:root {
  --light-gray: #f9f9f9;
  --black: #222222;
  --bright-blue: #3242FF;
  --check-green: #31BD70;
  --select-green: #E1F9EF;
}

#app {
  font-family: Arial, Helvetica, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: var(--black);
}
</style>
