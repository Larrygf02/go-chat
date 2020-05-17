const socket = io();

new Vue({
    el: "#chat-app",
    created() {
        socket.on("reply all", msg => {
            console.log(msg);
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