<template>
  <div class="main">
    <h1>{{ info }}</h1>
  </div>
</template>

<script>
  export default {
    el: '#main',
    data() {
      return {
        info: null
      }
    },
    mounted() {
      axios
        .get('https://api.coindesk.com/v1/bpi/currentprice.json')
        .then(response => (this.info = response))
    }
  }
</script>
