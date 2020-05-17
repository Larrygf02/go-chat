const socket = io();

new Vue({
    el: "#chat-app",
    created() {

    },
    data: {
        message: '',
        messages: []
    },
    methods: {
        sendMessage() {
            console.log(socket);
            socket.emit('chat message', this.message)
            this.message = ''
        }
    }
})