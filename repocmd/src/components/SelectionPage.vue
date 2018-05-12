<template>
<div class="selection-page">
  <div class="selection-header">
    <h3>Selected Repositories ({{selectedLen}})</h3>
    <!-- <p class="selection-links">
      <i data-feather="delete"></i><a href="#">Clear Selection</a>
    </p> -->
  </div>
  <div class="wrapper">
    <selection-list :repos="repos" :height="listHeight" @unchecked="unchecked" />
    <selection-actions :repos="repos"/>
  </div>
</div>
</template>

<script>
import SelectionList from './SelectionList'
import SelectionActions from './SelectionActions'

export default {
  name: 'SelectionPage',
  components: {
    SelectionList,
    SelectionActions
  },
  props: ['repos'],
  data () {
    return {
      listHeight: 300
    }
  },
  computed: {
    selectedLen () {
      if (this.repos) {
        return Object.keys(this.repos).length
      }
      return 0
    }
  },
  methods: {
    getMarginHeights (el) {
      const styles = window.getComputedStyle(el)
      return parseFloat(styles.marginTop) + parseFloat(styles.marginBottom)
    },
    unchecked (repo) {
      this.$emit('unchecked', repo)
    },
    calculateListHeight () {
      const winHeight = document.documentElement.clientHeight

      const header = document.querySelector('.selection-header')
      let headerHeight = 0
      if (header) {
        headerHeight = header.offsetHeight + this.getMarginHeights(header)
      }

      const actions = document.querySelector('.selection-actions')
      let actionsHeight = 0
      if (actions) {
        actionsHeight = actions.offsetHeight + this.getMarginHeights(actions)
      }

      // Find out how much padding the list has.
      const list = document.querySelector('.selection-list')
      let listPadding = 0
      if (list) {
        const listStyles = window.getComputedStyle(list)
        listPadding = parseFloat(listStyles.paddingTop) + parseFloat(listStyles.paddingBottom)
      }

      this.listHeight = winHeight - (headerHeight + actionsHeight) - listPadding
    }
  },
  mounted () {
    // window.addEventListener('resize', this.calculateListHeight)
    // this.calculateListHeight()
  },
  beforeDestroy () {
    // window.removeEventListener('resize', this.calculateListHeight)
  }
}
</script>

<style lang="scss" scoped>
.selection-page {
  padding-right: 20px;
}
.wrapper {
  display: grid;
  grid-template-columns: minmax(400px, auto) minmax(200px, 300px);
  grid-column-gap: 20px;
}
.selection-header {
  overflow: auto;
  padding: 0;
  text-align: left;
  h3 {
    font-size: 20px;
  }
  .selection-links {
    .feather {
      vertical-align: middle;
      height: 18px;
    }
    a {
      color: var(--bright-blue);
      font-weight: bold;
      text-decoration: none;
      vertical-align: middle;
      margin-left: 3px;
    }
  }
}
</style>
