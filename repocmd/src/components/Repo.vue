<template>
  <div class="repo">
    <h3 class="title"><a :href="repo.url">{{repo.name}}</a> <span v-if="repo.isPrivate" class="private-badge">Private</span></h3>
    <p class="description">{{repo.description}}</p>
    <p class="updated-at">Last updated: {{repo.updatedAt | moment("from", "now")}}</p>
 </div>
</template>

<script>
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
