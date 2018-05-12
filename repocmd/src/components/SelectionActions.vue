<template>
<div class="selection-actions">
  <form class="selection-summary">
    <fieldset>
      <legend>Selection Summary</legend>
      <div class="chart-wrapper">
        <bar-chart :chartData="stats"/>
      </div>
    </fieldset>
    </form>
  <form class="danger-zone">
    <fieldset>
      <legend>Danger Zone</legend>
      <select v-model="selected">
        <optgroup>
          <option value="">Select an action...</option>
          <option value="archive">Archive</option>
        </optgroup>
      </select>
      <div class="action-info">
        <p v-if="selected == 'archive'">
          Archiving can only be reversed through the GitHub website.
        </p>
        <p v-else>
          Some of these actions are irreversable. Make sure you understand the potential  side-effects before doing them.
        </p>
      </div>
      <div class="action-btn-wrapper">
        <button type="submit" class="action-btn" :disabled="!selected" href="#">I understand, do it</button>
      </div>
    </fieldset>
  </form>
</div>
</template>

<script>
import BarChart from './BarChart'
import feather from 'feather-icons'

export default {
  components: {
    BarChart
  },
  props: ['repos'],
  data () {
    return {
      selected: '',
      dropdownOpts: [{name: 'Archive', value: 'archive'}],
      stats: {}
    }
  },
  watch: {
    repos: {
      immediate: true,
      deep: true,
      handler (repos) {
        const stats = {
          watchers: 0,
          forks: 0,
          stars: 0
        }
        for (let k in repos) {
          const repo = repos[k]
          stats.watchers += repo.watchers.totalCount
          stats.forks += repo.forks.totalCount
          stats.stars += repo.stargazers.totalCount
        }
        this.stats = {
          labels: ['Watchers', 'Forks', 'Stars'],
          datasets: [
            {
              label: 'GitHub Stats',
              backgroundColor: '#3242FF',
              data: [stats.watchers, stats.forks, stats.stars]
            }
          ]
        }
      }
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
    text-align: center;
    padding: 0;

    .chart-wrapper {
      max-width: 250px;
    }

    .selection-summary {
      margin-top: -10px;
    }

    form {
      margin-bottom: 20px;
      fieldset {
        border: 1px solid #DDD;
        padding: 10px 20px;
      }
      legend {
        font-size: 16px;
        margin: 0;
      }
    }

    select {
      margin: 20px auto;
    }
    .action-info {
      font-size: 14px;
      background: #F9F9F9;
      width: 100%;
      height: 100px;
      max-height: 300px;
      margin: 0 auto;
      margin-bottom: 20px;
      display: flex;
      align-items: center;
      p {
        padding: 0 10px;
        line-height: 1.3;
      }
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
        cursor: pointer;
        &:hover {
          border: 1px solid black;
          color: black;
        }

        &[disabled] {
          border: 1px solid #dddddd;
          color: #ddd;
          pointer-events: none;
          cursor: default;
        }
      }
    }
  }
</style>
