<template>
  <div>
    <h3>Selected Repositories ({{repoCount}})</h3>
    <div class="select-list">
      <select-item v-for="repo of repos" v-bind:key="repo.id" :repo="repo"></select-item>
    </div>
    <select-controls></select-controls>
  </div>
</template>

<script>
import SelectItem from './SelectItem'
import SelectControls from './SelectControls'

export default {
  name: 'SelectList',
  components: {
    SelectItem,
    SelectControls
  },
  props: ['repos'],
  computed: {
    repoCount () {
      if (this.repos) {
        return Object.keys(this.repos).length
      }
      return 0
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .select-list {
    height: 500px;
    max-height: 500px;
    overflow-y: scroll;
    background: orange;
  }
  h3 {
    text-align: center;
    font-size: 22px;
    margin: 30px 0;
  }
</style>
