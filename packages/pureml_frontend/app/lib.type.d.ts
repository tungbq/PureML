declare module "marked";

export type dataset = {
  id: string;
  name: string;
  updated_at: string;
  created_by: string;
  uuid: string;
  updated_by: string;
  updated_by_name: string;
};

export type model = {
  id: string;
  name: string;
  updated_at: string;
  uuid: string;
  created_by: string;
  updated_by: string;
  updated_by_name: string;
};

export type org = {
  uuid: string;
  id: string;
  name: string;
};
