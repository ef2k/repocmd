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
      <repo v-for="repo in repos" v-bind:key="repo.id" :repo="repo"></repo>
    </div>
 </div>
</template>

<script>
import Repo from './Repo'

export default {
  components: {
    Repo
  },
  data () {
    return {
      loading: false,
      repos: []
    }
  },
  computed: {
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
      this.$http.get('http://localhost:5000/repos')
        .then(response => {
          this.repos = response.data
          this.loading = false
        })
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
</style>
