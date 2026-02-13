import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import Backend from '../lib/api'
import { Handle, Position } from "@xyflow/react";


type CompanyNodeProps = {
  id: string
  data: {
    label: string
  }
  type: "CompanyNode"
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

const CompanyForm = z.object({
  name: z.string(),
});

type CompanyFormType = z.infer<typeof CompanyForm>;

type NewCompany = {
  ID: string
  name: string
  node_id: string
  CreatedAt: string 
  UpdatedAt: string 
  DeletedAt: string
}

export function CompanyNode(props: CompanyNodeProps) {
  console.log(props)
  const { register, handleSubmit } = useForm<CompanyFormType>({
    resolver: zodResolver(CompanyForm),
  });

  const onSubmit = async (data: CompanyFormType) => {
    const createCompany = Backend.create().
      companies().
      send<CompanyFormType & {node_id: string}, NewCompany>({...data, node_id: props.id});
    const createNode = Backend.create().
      nodes().
      send<CompanyNodeProps, CompanyNodeProps>(props);

    const results = await Promise.allSettled([createCompany, createNode])

    if(results[0].status === "rejected") {
      alert("couldn't save company name"); 
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
            placeholder="Company Name..."
          />
          <button type="submit" className="btn btn-sm">Create</button>
        </label>
      </form>
      <Handle type="source" position={Position.Right} />
    </div>
  );
}
