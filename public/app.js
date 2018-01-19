new Vue({
    el: '#app',

    data: {
        ws: null, 
        amount: null, 
        adress: null,
       
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
