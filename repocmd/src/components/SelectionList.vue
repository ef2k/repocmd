<template>
  <div class="selection-list" :style="{ height: listHeight+'px', maxHeight: listHeight+'px'}">
    <selection v-for="repo of repos" v-bind:key="repo.id" :repo="repo" @unchecked="unchecked"/>
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
  data () {
    return {
      listHeight: document.documentElement.clientHeight
    }
  },
  mounted () {
    window.addEventListener('resize', this.handleResize)
    this.handleResize()
  },
  beforeDestroy () {
    window.removeEventListener('resize', this.handleResize)
  },
  methods: {
    handleResize () {
      this.listHeight = document.documentElement.clientHeight
    },
    unchecked (repo) {
      this.$emit('unchecked', repo)
    }
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
