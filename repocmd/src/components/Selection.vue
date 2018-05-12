<template>
  <div class="selection">
    <div class="icon-wrapper">
      <i data-feather="github"></i>
    </div>
    <div class="select-details">
      <h4>{{repo.name}}</h4>
      <p class="description">{{repo.description}}</p>
      <p class="date">Last commit: <strong>{{repo.pushedAt | moment("MMMM Do, YYYY")}}</strong></p>
    </div>
    <div class="icon-remove">
      <span @click="remove(repo)">
      <i data-feather="x-circle"></i>
      </span>
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
  },
  methods: {
    remove (repo) {
      this.$emit('unchecked', repo)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  .selection {
    margin: 0 0 10px 0;
    display: grid;
    grid-template-columns: 20px auto 25px;
    grid-column-gap: 12px;
    padding: 12px 14px;
    background: #f9f9f9;
  }
  .icon-wrapper {
    width: 20px;
    text-align: left;
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
    font-size: 16px;
  }
  .select-details .description {
    font-size: 14px;
    margin-bottom: 5px;
  }
  .select-details .date {
    font-size: 14px;
  }
  .icon-remove {
    .feather {
      color: #c0c0c0;
      transition: color 0.3s;
    }
    .feather:hover {
      color: red;
      cursor: pointer;
    }
  }
</style>
