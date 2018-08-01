<script>
new Vue({
    el: 'body',

    data: {
        tasks: [],
        newTask: {}
    },

    // This is run whenever the page is loaded to make sure we have a current task list
    created: function() {
        // Use the vue-resource $http client to fetch data from the /tasks route
        this.$http.get('/tasks').then(function(response) {
            this.tasks = response.data ? response.data : []
        })
    },

    methods: {
        createTask: function() {
            if (!$.trim(this.newTask.name)) {
                this.newTask = {}
                return
            }

            // Post the new task to the /tasks route using the $http client
            this.$http.put('/tasks', this.newTask).success(function(response) {
                this.newTask.id = response.created
                this.tasks.push(this.newTask)
                console.log("Task created!")
                console.log(this.newTask)
                this.newTask = {}
            }).error(function(error) {
                console.log(error)
            });
        },

        deleteTask: function(index) {
            // Use the $http client to delete a task by its id
            this.$http.delete('/tasks/' + this.tasks[index].id).success(function(response) {
                this.tasks.splice(index, 1)
                console.log("Task deleted!")
            }).error(function(error) {
                console.log(error)
            })
        }
    }
})
</script>