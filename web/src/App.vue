<script setup>

    import { ref, onMounted } from 'vue'
    const itemsList = ref([])

    const itemName = ref('')
    const itemDescription = ref('')

    const myRequest = new Request('david-pw.com')

    function getItems(event) {
        //itemsList.value.push(itemName.value + itemDescription.value)

        fetch('http://127.0.0.1:8080/items/', {method: 'GET'}).then((response) => {    
            console.log("Hello")
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status})`)
            }
            response.body

            itemsList.value.push(response.body)

            return response.blob()
        })

    }

    onMounted(() => {
    })

</script>

<template>
    <header>

        <div class="wrapper">
        </div>
    </header>
    <main>
        <h1>Test List</h1>


        <ul>
            <li v-for="item in itemsList">
                {{ item }}
            </li>
        </ul>

        <input v-model="itemName" placeholder="Item Name"> 
        <input v-model="itemDescription" placeholder="Item Description">

        <button @click="getItems">Get Items</button>

    </main>
</template>

<style scoped></style>
