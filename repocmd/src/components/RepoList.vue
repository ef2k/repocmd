<template>
<div class="repo-list">
  <div class="side-pane" :style="{ height: winHeight+'px', maxHeight: winHeight+'px'}">
    <div v-if="loading">
      <p><i data-feather="github"></i> Fetching the latest repo data...</p>
    </div>
    <div v-else-if="error">
      <p v-if="error.status === 0">The proxy server is offline.</p>
      <p v-else>There was a problem connecting to the GitHub API.</p>
      <pre>
        {{error}}
      </pre>
      <a href="#">Retry the server</a>
    </div>
    <div v-else>
      <div>
        <h2>Repositories ({{repoLen}})</h2>
        <p class="filter-options">
          Filter by:
          <a href="#" @click="filterBy('all')" :disabled="isFilteredAll">All</a> |
          <a href="#" @click="filterBy('public')" :disabled="isFilteredPublic">Public</a> |
          <a href="#" @click="filterBy('private')" :disabled="isFilteredPrivate">Private</a>
        </p>
        <repo v-for="repo in filteredList" v-bind:key="repo.id" :repo="repo" @checked="checkRepo" @unchecked="uncheckRepo"></repo>
      </div>
    </div>
  </div>
  <div class="main-pane" :style="{ height: winHeight+'px', maxHeight: winHeight+'px'}">
    <transition name="fade">
      <selection-page v-if="hasSelected" :repos="selected"/>
    </transition>
  </div>
</div>
</template>

<script>
import Repo from './Repo'
import SelectionPage from './SelectionPage'

import feather from 'feather-icons'

export default {
  components: {
    Repo,
    SelectionPage
  },
  data () {
    return {
      loading: false,
      error: '',
      filterOption: 'all',
      repos: [],
      selected: {},
      winHeight: document.documentElement.clientHeight
    }
  },
  computed: {
    hasSelected () {
      return Object.keys(this.selected).length > 0
    },
    repoLen () {
      if (this.filteredList) {
        return this.filteredList.length
      }
      return ''
    },
    isFilteredAll () {
      return this.filterOption === 'all'
    },
    isFilteredPublic () {
      return this.filterOption === 'public'
    },
    isFilteredPrivate () {
      return this.filterOption === 'private'
    },
    filteredList () {
      if (this.filterOption === 'private') {
        return this.repos.filter(repo => repo.isPrivate)
      } else if (this.filterOption === 'public') {
        return this.repos.filter(repo => !repo.isPrivate)
      }
      return this.repos
    }
  },
  created () {
    this.fetchRepos()
  },
  mounted () {
    feather.replace()
    window.addEventListener('resize', this.handleResize)
  },
  beforeDestroy () {
    window.removeEventListener('resize', this.handleResize)
  },
  methods: {
    fetchRepos () {
      this.loading = true
      this.$http.get('http://localhost:3000/repos')
        .then(response => {
          this.repos = response.data
          this.loading = false
        })
        .catch(err => {
          this.loading = false
          this.error = err
        })
    },
    checkRepo (repo) {
      const item = this.repos.find(item => item.id === repo.id)
      this.$set(item, 'checked', true)
      this.addToSelectList(repo)
    },
    uncheckRepo (repo) {
      const item = this.repos.find(item => item.id === repo.id)
      this.$set(item, 'checked', false)
      this.removeFromSelectList(repo)
    },
    addToSelectList (repo) {
      this.$set(this.selected, repo.id, repo)
    },
    removeFromSelectList (repo) {
      this.$delete(this.selected, repo.id)
    },
    filterBy (option) {
      this.filterOption = option
    },
    handleResize () {
      this.winHeight = document.documentElement.clientHeight
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  .repo-list {
    display: grid;
    grid-template-columns: 470px auto;
    grid-column-gap: 20px;
  }
  .side-pane {
    border-right: 1px solid #D8D8D8;
    overflow-y: scroll;
  }
  .main-pane {
    position: relative;
  }
  .filter-options a {
    color: var(--bright-blue);
    text-decoration: none;
    font-weight: bold;
  }
  .filter-options a[disabled=disabled] {
    color: var(--black);
  }
</style>
