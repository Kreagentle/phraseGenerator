<template>
  <div id="app">
    <ul>
    <p1>Text.Text.Text.</p1>
    <a href="https://vk.com/christina_belyavskaya" target="_blank"><img src="./assets/vk1.png"></a>
    <a href="https://www.instagram.com/bellyavskaya/" target="_blank"><img src="./assets/instagram1.png"></a>
    </ul>
    <p2>{{ msg }}</p2>
    <p2>{{info}}</p2>
    <ul>
      <li><a href="https://vuejs.org" target="_blank">Core Docs</a></li>
      <li><a href="https://forum.vuejs.org" target="_blank">Forum</a></li>
    </ul>
  </div>
</template>

<script>

  import axios from 'axios'
export default {
  name: 'app',
  data () {
    return {
      msg: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Quis ipsum suspendisse ultrices gravida. Risus commodo viverra maecenas accumsan lacus vel facilisis.',
      info: null
    }
  },

  created() {
    this.fetch()
  },

  methods: {
    fetch() {
      axios
        .get('/api/phrases')
        .then(response => (this.info = response));
    }
  },
}
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

p1, p2 {
  font-weight: normal;
  margin: 0 10px;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

</style>
