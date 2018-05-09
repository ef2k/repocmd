<template>
<div class="selection-actions">
  <div>
    <h4>Apply an action</h4>
    <select v-model="selected">
      <optgroup>
        <option disabled value="">Select an action...</option>
        <option value="archive">Archive</option>
      </optgroup>
    </select>
    <div class="action-info">
      <p v-if="selected == 'archive'">
        <i data-feather="info"></i> Archiving can only be reversed through the GitHub website.
      </p>
    </div>
    <div class="action-btn-wrapper">
      <a class="action-btn" href="#">Apply the action</a>
    </div>
  </div>
</div>

</template>

<script>
import feather from 'feather-icons'
export default {
  components: {
  },
  data () {
    return {
      selected: '',
      dropdownOpts: [{name: 'Archive', value: 'archive'}]
    }
  },
  mounted () {
    feather.replace()
  },
  methods: {
    select (payload) {
      this.selected = payload.value
    }
  }
}
</script>

<style lang="scss" scoped>
  .selection-actions {
    position: absolute;
    left: 0;
    bottom: 0;
    width: 100%;
    border-top: 1px solid #dddddd;
    text-align: center;
    padding: 20px 0;

    h4 {
      font-size: 20px;
      margin: 0;
    }
    select {
      margin: 20px auto;
    }
    .action-info {
      background: #F9F9F9;
      width: 300px;
      height: 100px;
      max-height: 300px;
      margin: 0 auto;
      margin-bottom: 20px;
      display: flex;
      align-items: center;
    }
    .action-btn-wrapper {
      display: block;
      margin-bottom: 10px;
      .action-btn {
        border: 1px solid var(--bright-blue);
        font-weight: 600;
        font-size: 14px;
        line-height: 14px;
        color: var(--bright-blue);
        border-radius: 3px;
        text-decoration: none;
        padding: 12px 15px;
        display: inline-block;
        transition: all 0.4s;
        &:hover {
          border: 1px solid black;
          color: black;
        }
      }
    }
  }
</style>
