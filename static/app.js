const app = Vue.createApp({
  data() {
    return {
      entries: [],
      content: '',
    }
  },
  created() { this.update() },
  methods: {
    del: (e) => {
      const id = e.target.id.substring(5);
      axios.delete('/entries/' + id)
        .then(() => {
          app.content = ''
          app.update()
        })
        .catch((error) => console.log(error));
    },
    add: () => {
      const payload = {'content': app.content}
      axios.post('/entries', payload)
        .then(() => {
          app.content = ''
          app.update()
        })
        .catch((err) => {
          alert(err.response.data || err.message)
        })
    },
    update: () => {
      axios.get('/entries')
        .then((response) => app.entries = response.data || [])
        .catch((error) => console.log(error));
    }
  }
}).mount('#app');
