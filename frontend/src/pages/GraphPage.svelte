<script lang="ts">
    import { onMount } from "svelte";
    import { GetEdgeList, GetLookUp } from "../../wailsjs/go/main/App";
    import type {repo} from '../../wailsjs/go/models'
    
    let graph_container;
    const graph_nodes = {};
    let edgeList: {
        [key: string]: repo.Edge[];
    };
    let lookUp: {
        [key: string]: string;
    };

    onMount(async () => {
        const edgeListPromise = GetEdgeList();
        const lookUpPromise = GetLookUp();

        edgeList = await edgeListPromise;
        lookUp = await lookUpPromise;

        for(const key in lookUp) {

        }

        
        //@ts-ignore
        let graphJSON = {
        "nodes": ["mark", "higgs", "other", "etc"],
        "edges": [
            ["mark", "higgs"],
            ["mark", "etc"],
            ["mark", "other"]
        ]
        };
        
        //@ts-ignore
        let graph = new Springy.Graph();
        graph.loadJSON(graphJSON);
        //@ts-ignore
        var layout = new Springy.Layout.ForceDirected(graph, 400.0, 400.0, 0.5);
        //@ts-ignore
        var renderer = new Springy.Renderer(layout,
            function clear() {
                // code to clear screen
            },
            function drawEdge(edge, p1, p2) {
                // draw an edge
            },
            function drawNode(node, p) {
                // draw a node
            }
        );
        renderer.start();
    })
</script>

<main>
    <div>Hello</div>
    <div>
        {JSON.stringify(edgeList)}
    </div>

    <div bind:this={graph_container}></div>
</main>

<style></style>