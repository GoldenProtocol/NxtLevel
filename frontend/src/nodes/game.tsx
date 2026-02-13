import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import Backend from '../lib/api'
import { Handle, Position } from "@xyflow/react";


type GameNodeProps = {
  id: string
  data: {
    label: string
  }
  type: "GameNode"
  positionAbsoluteX: number
  positionAbsoluteY: number
  selected: boolean
  selectable: boolean
  draggable: boolean
  deletable: boolean
  isConnectable: boolean
  dragging: boolean
  zIndex: number
  width: number
  height: number
}

const GameForm = z.object({
  name: z.string(),
});

type GameFormType = z.infer<typeof GameForm>;

type NewCompany = {
  ID: string
  name: string
  node_id: string
  CreatedAt: string 
  UpdatedAt: string 
  DeletedAt: string
}

export function GameNode(props: GameNodeProps) {
  const { register, handleSubmit } = useForm<GameFormType>({
    resolver: zodResolver(GameForm),
  });

  const onSubmit = async (data: GameFormType) => {
    const createGame = Backend.create().
      games().
      send<GameFormType & {node_id: string}, NewCompany>({...data, node_id: props.id});
    const createNode = Backend.create().
      nodes().
      send<GameNodeProps, GameNodeProps>(props);

    const results = await Promise.allSettled([createGame, createNode])

    if(results[0].status === "rejected") {
      alert("couldn't save game name"); 
    }

    if (results[1].status === "rejected") {
      alert("couldn't save node")
    }
  };

  return (
    <div>
        <form onSubmit={handleSubmit(onSubmit)}>
      <label className="input">
        <input
          {...register("name")}
          type="text"
          className="grow"
          placeholder="Game Name..."
        />
        <button type="submit" className="btn btn-sm">Create</button>
      </label>
    </form>
    <Handle type="target" position={Position.Left} />
    </div>
  );
}
