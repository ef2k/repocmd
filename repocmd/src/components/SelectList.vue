<template>
  <div class="select-list">
    <h3>Selected Repos</h3>
    <ul>
      <li v-for="repo of repos" v-bind:key="repo.id">
        {{repo.name}}, last updated {{repo.pushedAt | moment("from", "now")}}
       </li>
    </ul>
  </div>
</template>

<script>

export default {
  name: 'SelectList',
  props: ['repos']
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .select-list {
    background: black;
    color: white;
    min-width: 200px;
    min-height: 100px;
    position: fixed;
    bottom: 10px;
    right: 10px;
    padding: 20px;
  }
</style>
