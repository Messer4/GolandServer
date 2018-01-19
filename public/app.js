new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
      //  newMsg: '', // Holds new messages to be sent to the server
       // chatContent: '', // A running list of chat messages displayed on the screen
        amount: null, // Email address used for grabbing an avatar
        adress: null, // Our username
       // joined: false // True if email and username have been filled in
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
           
        });
    },

    methods: {
        send: function () {
            if (!this.amount) {
                Materialize.toast('You must enter an email', 2000);
                return
            }
            if (!this.adress) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            
                this.ws.send(
                    JSON.stringify({
                        amount: $('<p>').html(this.amount).text(),
                        adress: this.username = $('<p>').html(this.adress).text(),
                       
                    }
                ));
               
            
        },

       

       
    }
});
