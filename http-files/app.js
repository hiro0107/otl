const App = {
    data() {
        return {
        counter: 0
        }
    },
    template: "<log-portal/>"
}
const app = Vue.createApp(App)
app.component('log-portal', {
    template: "<div>ok</div>"
})
 
app.mount('#app')
