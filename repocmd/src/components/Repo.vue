<template>
  <div class="repo">
    <h3 class="title">
      <i data-feather="github"></i>
      <a :href="repo.url">{{repo.name}}</a>
      <span v-if="repo.isPrivate" class="mini-badge">Private</span>
      <span v-if="repo.isFork" class="mini-badge">Fork</span>
    </h3>
    <p class="description">{{repo.description}}</p>
    <p>
      <span class="repo-stat">
        <i data-feather="eye"></i> Watchers: {{repo.watchers.totalCount}}
      </span>
      <span class="repo-stat">
        <i data-feather="star"></i> Stars: {{repo.stargazers.totalCount}}
      </span>
      <span class="repo-stat">
        <i data-feather="git-branch"></i> Forks: {{repo.forks.totalCount}}
      </span>
    </p>
    <p class="updated-at">Last update: {{repo.updatedAt | moment("from", "now")}}</p>
    <p class="pushed-at">Last commit: <i data-feather="git-commit"></i> {{repo.pushedAt | moment("from", "now")}}</p>
 </div>
</template>

<script>
import feather from 'feather-icons'
export default {
  props: ['repo'],
  methods: {
    archive (repo, event) {
      console.log('The event is ', repo, event)
      repo.isLoading = true
      window.setTimeout(1000, () => {
        repo.isLoading = false
      })
    }
  },
  created () {
    feather.replace()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .feather {
    width: 14px;
    height: 14px;
    line-height: 14px;
    display: inline-block;
    vertical-align: middle;
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
  .mini-badge {
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
  .repo-stat {
    margin-right: 12px;
  }
  .description {
    line-height: 1.5;
  }
</style>
