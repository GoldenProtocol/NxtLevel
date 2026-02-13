import { useCallback, useEffect } from "react";
import {
  addEdge,
  Background,
  Controls,
  ReactFlow,
  ReactFlowProvider,
  useEdgesState,
  useNodesState,
} from "@xyflow/react";
import { nodeTypes } from "./nodes";
import "@xyflow/react/dist/style.css";
import { Sidebar } from "./sidebar";
import Backend from './lib/api'

const initialNodes = [
  {
    id: "1",
    type: "CompanyNode",
    data: { label: "input node" },
    position: { x: 250, y: 5 },
  },
];

function DnDFlow() {
  const [nodes, _, onNodesChange] = useNodesState(initialNodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);

  console.log(edges)

  const onConnect = useCallback(
    (params: Connection) => setEdges((eds) => addEdge(params, eds)),
    [],
  );

  useEffect(() => {
    async function storeEdges(edges: any) {
      await Backend.connections().create().send(edges)
    }

    if (edges.length > 0) {
      storeEdges(edges)
    }
  }, [edges.length])

  return (
    <div className="dndflow">
      <div className="reactflow-wrapper">
        <ReactFlow
          nodes={nodes}
          edges={edges}
          onNodesChange={onNodesChange}
          onEdgesChange={onEdgesChange}
          nodeTypes={nodeTypes}
          onConnect={onConnect}
          fitView
        >
          <Controls />
          <Background />
        </ReactFlow>
      </div>
      <Sidebar />
    </div>
  );
}

export default () => (
  <ReactFlowProvider>
    <DnDFlow />
  </ReactFlowProvider>
);
