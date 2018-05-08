<template>
  <div id="app">
    <div class="navbar">

    </div>
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

html, body {
  padding: 0;
  margin: 0;
}

#app {
  font-family: Arial, Helvetica, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: var(--black);
  display: grid;
  grid-template-columns: 50px auto;
  grid-column-gap: 0px;
}
.navbar {
  background: black;
}
</style>
