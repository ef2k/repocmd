<template>
  <div class="selection-list" :style="{ height: height+'px', maxHeight: height+'px'}">
    <selection v-for="repo of repos" v-bind:key="repo.id" :repo="repo"/>
  </div>
</template>

<script>
import Selection from './Selection'

export default {
  components: {
    Selection
  },
  props: [
    'repos',
    'height'
  ],
  computed: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .selection-list {
    overflow: auto;
  }
  h3 {
    text-align: center;
    font-size: 22px;
    margin: 20px 0;
  }
</style>
