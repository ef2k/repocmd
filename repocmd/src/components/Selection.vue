<template>
  <div class="selection">
    <div class="icon-wrapper">
      <i data-feather="github"></i>
    </div>
    <div class="select-details">
      <h4>{{repo.name}}</h4>
      <p class="description">{{repo.description}}</p>
      <p class="date">Last committed on <strong>{{repo.pushedAt | moment("MMMM Do, YYYY")}}</strong></p>
    </div>
  </div>
</template>

<script>
import feather from 'feather-icons'

export default {
  name: 'SelectItem',
  props: ['repo'],
  mounted () {
    feather.replace()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .selection {
    margin: 0 0 30px 0;
    display: grid;
    grid-template-columns: 40px auto;
    grid-column-gap: 5px;
  }
  .icon-wrapper {
    width: 40px;
    text-align: center;
  }
  .icon-wrapper .feather {
    margin-top: 2px;
    width: 20px;
    height: 20px;
  }
  .select-details * {
    margin: 0;
    padding: 0;
  }
  .select-details h4 {
    margin-bottom: 5px;
  }
  .select-details .description {
    margin-bottom: 5px;
  }
  .select-details .date {
    font-size: 14px;
  }
</style>
