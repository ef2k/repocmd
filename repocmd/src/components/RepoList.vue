<template>
<div class="repo-list">
  <div class="sidebar">
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
        <h2>Owned Repositories ({{repoLen}})</h2>
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

  <div class="repo-summary">
    <transition name="fade">
      <div v-if="hasSelected">
        <select-list :repos="selected"></select-list>
      </div>
    </transition>
  </div>
  </div>
</template>

<script>
import Repo from './Repo'
import SelectList from './SelectList'
import feather from 'feather-icons'

export default {
  components: {
    Repo,
    SelectList
  },
  data () {
    return {
      loading: false,
      error: '',
      filterOption: 'all',
      repos: [],
      selected: {}
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
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

  .repo-list {
    display: grid;
    grid-template-columns: 475px auto;
    grid-column-gap: 0px;
  }
  .sidebar {
    padding: 10px;
    border-right: 1px solid #D8D8D8;
    max-height: 800px;
    overflow-y: scroll;
  }
  .repo-summary {
    padding: 10px;
  }

  .repo {
    background: var(--light-gray);
    padding: 5px 30px;
    margin: 20px auto;
    display: block;
    border: 1px solid rgba(27, 30, 34, 0.04);
  }
  .title {
    font-size: 22px;
  }
  .title a {
    text-decoration: none;
    color: #3242FF;
  }
  .private-badge {
    margin-left: 5px;
    font-size: 12px;
    line-height: 11px;
    padding: 3px 4px;
    border: 1px solid rgba(27, 30, 34, 0.15);
    box-shadow: none;
    color: var(--black);
    display: inline-block;
    vertical-align: middle;
  }
  .description {
    line-height: 1.5;
  }

  .filter-options a {
    color: var(--bright-blue);
    text-decoration: none;
    font-weight: bold;
  }
  .filter-options a[disabled=disabled] {
    color: var(--black);
  }

  /* Fade transitions */
  .fade-enter-active, .fade-leave-active {
    transition: opacity .5s;
  }
  .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
    opacity: 0;
  }
</style>
