const appcomponent = {
    data() {
        return {
            query: null,
            data: []
        }
    },
    methods: {
        async getData() {
            await axios.get('https://api.adviceslip.com/advice/search/{query}${this.query}')
                .then((response) => {
                    this.data = response.data.slips
                    console.log(this.query)
                })
                .catch((error) => {
                    console.log(error)
                })
        }
    }
}
const app = Vue.createApp(appcomponent).mount('#app')