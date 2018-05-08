<template>
  <div class="container">
    <div v-if="loading">
      <p>Loading...</p>
    </div>
    <div v-else>
      <h2>Repositories ({{repoLen}})</h2>
      <p>
        Filter
        <a href="#">All</a> |
        <a href="#">Public</a> |
        <a href="#">Private</a>
      </p>
      <repo v-for="repo in repos" v-bind:key="repo.id" :repo="repo" v-on:checked="addToSelectList" v-on:unchecked="removeFromSelectList"></repo>
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

export default {
  components: {
    Repo,
    SelectList
  },
  data () {
    return {
      loading: false,
      repos: [],
      selected: {}
    }
  },
  computed: {
    hasSelected () {
      return Object.keys(this.selected).length > 0
    },
    repoLen () {
      if (this.repos) {
        return this.repos.length
      }
      return ''
    }
  },
  created () {
    this.fetchRepos()
  },
  methods: {
    fetchRepos () {
      this.loading = true
      this.$http.get('http://localhost:3000/repos')
        .then(response => {
          this.repos = response.data
          this.loading = false
        })
    },
    addToSelectList (repo) {
      // this.selected[repo.id] = repo
      this.$set(this.selected, repo.id, repo)
    },
    removeFromSelectList (repo) {
      // delete this.selected[repo.id]
      this.$delete(this.selected, repo.id)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
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

  /* Fade transitions */
  .fade-enter-active, .fade-leave-active {
    transition: opacity .5s;
  }
  .fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
    opacity: 0;
  }
</style>
