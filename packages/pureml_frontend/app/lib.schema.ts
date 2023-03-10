import { z } from 'zod';

export const versionDataSchema = z.string();
export type versionDataType = z.infer<typeof versionDataSchema>;

export const nodesSchema = z.array(
  z.object({
    id: z.string(),
    text: z.string(),
    nodeType: z.string().optional(),
  })
);
export type nodesType = z.infer<typeof nodesSchema>;

export const edgesSchema = z.array(
  z.object({ id: z.string(), from: z.string(), to: z.string() })
);
export type edgesType = z.infer<typeof edgesSchema>;
