<script lang="ts">
  import { onMount } from 'svelte';
  import cytoscape from 'cytoscape';
  import klay from 'cytoscape-klay';
  cytoscape.use(klay)
  let cy;

  const draw = () => {
    cy = cytoscape({container: document.getElementById('cy'),

    boxSelectionEnabled: false,
    autounselectify: true,

    style: cytoscape.stylesheet()
      .selector('node')
        .style({
          'content': 'data(id)',
          'color': 'white',
        })
      .selector('edge')
        .style({
          'curve-style': 'bezier',
          'target-arrow-shape': 'triangle',
          'width': 4,
          'line-color': '#ddd',
          'target-arrow-color': '#ddd'
        })
      .selector('.highlighted')
        .style({
          'background-color': '#61bffc',
          'line-color': '#61bffc',
          'target-arrow-color': '#61bffc',
          'transition-property': 'background-color, line-color, target-arrow-color',
          'transition-duration': '0.5s'
        }),

    elements: {
        nodes: [
          { data: { id: 'a' } },
          { data: { id: 'b' } },
          { data: { id: 'c' } },
          { data: { id: 'd' } },
          { data: { id: 'e' } }
        ],

        edges: [
          { data: { id: 'a"e', weight: 1, source: 'a', target: 'e' } },
          { data: { id: 'ab', weight: 3, source: 'a', target: 'b' } },
          { data: { id: 'be', weight: 4, source: 'b', target: 'e' } },
          { data: { id: 'bc', weight: 5, source: 'b', target: 'c' } },
          { data: { id: 'ce', weight: 6, source: 'c', target: 'e' } },
          { data: { id: 'cd', weight: 2, source: 'c', target: 'd' } },
          { data: { id: 'de', weight: 7, source: 'd', target: 'e' } }
        ]
      },

      layout: {
        name: 'breadthfirst',
        directed: true,
        roots: '#a',
        padding: 10
      }
    });
    
    //@ts-ignore
    window.cy = cy
    cy.center();
    cy.mount();

    document.querySelectorAll('canvas').forEach(cv => {
        cv.style.width = '100%';
        cv.style.height = '100%';
        cv.style.left = '0px';
        cv.style.top = '0px';
      });
    
    cy.getElementById('a').addClass('highlighted')
  }

  onMount(draw)
</script>

<main>
  <div id="cy" style="top:0px; left: 0px; height: 100vh; width: 100vw;"></div>
</main>

<style lang="scss">
  canvas {
    top:0px !important;
    left: 0px !important;
    height: 100% !important;
    width: 100% !important;
  }
</style> 