<template>
  <div class="main">
    <h1>{{ info }}</h1>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  el: '#main',
  data: function () {
    return {
      info: null
    }
  },
  mounted: function () {
    axios
      .get('https://api.coindesk.com/v1/bpi/currentprice.json')
      .then(response => (this.info = response))
  }
}
</script>a
