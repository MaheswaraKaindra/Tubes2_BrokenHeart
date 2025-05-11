'use client';
import React, { useMemo } from 'react';
import ReactFlow from 'reactflow';
import dagre from 'dagre';
import 'reactflow/dist/style.css';

// need to change to json bfore dipassing ke sini --> json.NewEncoder(w).Encode(tree)

const nodeWidth = 150;
const nodeHeight = 60;

const TreeVisualizer = ({ tree }) => {
  const treeToFlow = (tree) => {
    const dagreGraph = new dagre.graphlib.Graph();
    dagreGraph.setDefaultEdgeLabel(() => ({}));
    dagreGraph.setGraph({ rankdir: 'BT' });

    const nodes = [];
    const edges = [];

    const traverse = (node, parentId = null) => {
      if (!node) return;

      const id = `${nodes.length}`;
      nodes.push({
        id,
        data: { label: node.Name },
        position: { x: 0, y: 0 }, 
      });

      dagreGraph.setNode(id, { width: nodeWidth, height: nodeHeight });

      if (parentId !== null) {
        edges.push({
          id: `e${parentId}-${id}`,
          source: parentId,
          target: id,
          type: 'default',
        });
        dagreGraph.setEdge(parentId, id);
      }

      if (node.Left) traverse(node.Left, id);
      if (node.Right) traverse(node.Right, id);
    };

    traverse(tree);

    dagre.layout(dagreGraph);

    // Apply layouted positions
    nodes.forEach((node) => {
      const pos = dagreGraph.node(node.id);
      node.position = { x: pos.x, y: pos.y };
      node.sourcePosition = 'top';
      node.targetPosition = 'bottom';
    });

    return { nodes, edges };
  };

  const { nodes, edges } = useMemo(() => treeToFlow(tree), [tree]);

  return (
    <div style={{ width: '100%', height: '500px', border: '2px solid #ccc' }}>
      <ReactFlow
        nodes={nodes}
        edges={edges}
        fitView
        defaultEdgeOptions={{
          type: 'straight',
          style: { stroke: 'black' },
          markerEnd: {
            type: 'arrowclosed',
            color: 'black',
          },
        }}
        nodesDraggable={false}
        nodesConnectable={false}
        panOnDrag
        zoomOnScroll
      />
    </div>
  );
};

export default TreeVisualizer;
