const socket = io();

new Vue({
    el: "#chat-app",
    created() {
        socket.on("chat message", msg => {
            this.messages.push(msg)
        })
    },
    data: {
        message: '',
        messages: []
    },
    methods: {
        sendMessage() {
            socket.emit('chat message', this.message)
            this.message = ''
        }
    }
})