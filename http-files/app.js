const App = {
    template: "<log-portal/>"
}
const app = Vue.createApp(App)

const repeat = async (time, func) => {
    const f = async () => {
        await func()
        setTimeout(f, time)
    }
    await f()
}

app.component('log-portal', {
    template: "<div>ok</div>",
    setup(props) {
        const logs = Vue.reactive([])
        repeat(1000, async () => {
            const response = await fetch("/logs")
            const json = await response.json()
            console.log(json)
        })
        return {
            logs
        }
    }
})
 
app.mount('#app')
