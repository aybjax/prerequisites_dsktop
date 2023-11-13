<script lang="ts">
    import { DbConnection } from "../models/connection";
    import {ConnectNeo} from "../../wailsjs/go/main/App"
    import router from 'svelte-router';

    const db_connection = DbConnection.instance
    let url = db_connection.url
    let port = db_connection.port
    let user = db_connection.user
    let passwd = db_connection.passwd

    const checkConnection = () => {
        ConnectNeo({
            uri: url,
            port,
            user,
            pswd: passwd,
        }).then(_ => {
            console.log('connected')
            router.push('/graph')
        }).catch(e=> {
            alert('Not connected')
        })
    }
</script>

<main>
    <h1>Neo4j database connection</h1>

    <form action="">
        <div>
            <label for="url">url</label>
            <input id="url" bind:value={db_connection.url} />
        </div>
        <div>
            <label for="port">port</label>
            <input id="port" bind:value={db_connection.port} />
        </div>
        <div>
            <label for="user">user</label>
            <input id="user" bind:value={db_connection.user} />
        </div>
        <div>
            <label for="passwd">password</label>
            <input id="passwd" bind:value={db_connection.passwd} />
        </div>
        <button on:click|preventDefault={checkConnection}>Check connection</button>
    </form>
</main>

<style lang="scss">
    form {
        margin: {
            left: auto;
            right: auto;
        };
        max-width: 250px;
        display: flex;
        flex-direction: column;
    }
    label {
        display: block;
        margin: {
            bottom: 5px;
        };
        text-align: start;
    }
    input {
        width: 100%;
        margin: {
            bottom: 15px;
        };
    }
    button {
        width: 100px;
        align-self: center;
    }
</style>