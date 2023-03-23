import {
  Form,
  useActionData,
  useLoaderData,
  useSubmit,
} from "@remix-run/react";
import { ChevronDown, PlusCircle } from "lucide-react";
import { Suspense, useState } from "react";
import Tabbar from "~/components/Tabbar";
import AvatarIcon from "~/components/ui/Avatar";
import Button from "~/components/ui/Button";
import Loader from "~/components/ui/Loading";
import {
  fetchOrgDetails,
  updateOrgAddMember,
  updateOrgChangeRole,
  updateOrgRemoveMember,
} from "~/routes/api/org.server";
import { getSession } from "~/session";
import Modal from "~/components/ui/Modal";
import Input from "~/components/ui/Input";
import { toast } from "react-toastify";
import * as SelectPrimitive from "@radix-ui/react-select";

export async function loader({ request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const accesstoken = session.get("accessToken");
  const orgId = session.get("orgId");
  if (!orgId) return null;
  const orgDetails = await fetchOrgDetails(orgId, accesstoken);
  return orgDetails;
}

export const action = async ({ request }: any) => {
  const form = await request.formData();
  const orgId = form.get("orgid");
  const addEmail = form.get("addEmail");
  const removeEmail = form.get("removeEmail");
  const role = form.get("role");
  const roleEmail = form.get("roleEmail");
  const session = await getSession(request.headers.get("Cookie"));
  const accessToken = session.get("accessToken");
  let addMemberData;
  let removeMemberData;
  let changeRole;
  if (addEmail) {
    addMemberData = await updateOrgAddMember(orgId, addEmail, accessToken);
  } else addMemberData = null;
  if (removeEmail) {
    removeMemberData = await updateOrgRemoveMember(
      orgId,
      removeEmail,
      accessToken
    );
  } else removeMemberData = null;
  if (role) {
    changeRole = await updateOrgChangeRole(orgId, roleEmail, role, accessToken);
  } else changeRole = null;
  return { addMemberData, removeMemberData, changeRole };
};

const allRoles = [
  { value: "owner", label: "Owner" },
  { value: "member", label: "Member" },
];

export default function MembersSetting() {
  const data = useLoaderData();
  const updateMemberData = useActionData();
  const submit = useSubmit();
  const [role, setRole] = useState("");
  if (updateMemberData) {
    if (updateMemberData.addMemberData) {
      if (
        updateMemberData.addMemberData.message === "User added to organization"
      )
        toast.success("Member added to this organization!");
      else if (
        updateMemberData.addMemberData.message ===
        "User already added to organization"
      )
        toast.error("User already exists!");
      else if (
        updateMemberData.addMemberData.message === "User to add not found"
      )
        toast.error("User does not exist!");
      else toast.error("Something went wrong!");
    }
    if (updateMemberData.removeMemberData) {
      if (
        updateMemberData.removeMemberData.message ===
        "User removed from organization"
      )
        toast.success("Member removed from this organization!");
      else if (
        updateMemberData.removeMemberData.message === "User to remove not found"
      )
        toast.error("Member not found in this organization!");
      else if (
        updateMemberData.removeMemberData.message ===
        "Owner can't be removed from organization"
      )
        toast.error("Owner can't be removed from organization!");
      else toast.error("Something went wrong!");
    }
    if (updateMemberData.changeRole) {
      if (updateMemberData.changeRole.message === "User role updated")
        toast.success("Role updated!");
      else if (
        updateMemberData.changeRole.message ===
        "Role must be one of 'owner' or 'member'"
      )
        toast.error("Can't update this role!");
      else if (
        updateMemberData.changeRole.message ===
        "You are not authorized to update users in this organization"
      )
        toast.error("Only Owners can update the role of this organization!");
      else toast.error("Something went wrong!");
    }
  }
  // ##### dropdown role switch functionality #####
  function roleChange(event: any) {
    setRole(event.target.value);
    submit(event.currentTarget, { replace: true });
  }

  return (
    <Suspense fallback={<Loader />}>
      <div className="flex justify-center w-full border-b-2 border-slate-100">
        <div className="w-full 2xl:max-w-screen-2xl">
          <Tabbar intent="primarySettingTab" tab="members" />
        </div>
      </div>
      <div className="flex justify-center w-full">
        <div className="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl">
          <div className="pt-8 px-12">
            <div className="flex justify-between pb-10">
              <div className="font-medium">Organization Members</div>
              <Modal
                btnName={
                  <Button intent="primary" fullWidth={false}>
                    <PlusCircle className="text-white w-4" />
                    <div className="pl-2">Add member</div>
                  </Button>
                }
                title="Add Member"
              >
                <Form method="post" reloadDocument className="w-full">
                  <div>
                    <input
                      className="hidden"
                      type="text"
                      name="orgid"
                      defaultValue={data[0].uuid}
                    />
                    <label htmlFor="addEmail" className="text-sm pb-1">
                      Email
                      <Input
                        intent="valuePrimary"
                        type="text"
                        name="addEmail"
                        fullWidth={false}
                        defaultValue=""
                        aria-label="addEmail"
                        data-testid="addEmail"
                        required
                      />
                    </label>
                  </div>
                  <div className="pt-12 grid justify-items-end w-full">
                    <div className="flex justify-end w-1/2">
                      <Button
                        intent="primary"
                        type="submit"
                        className="pl-2"
                        fullWidth={false}
                      >
                        Add
                      </Button>
                    </div>
                  </div>
                </Form>
              </Modal>
            </div>
            <div className="overflow-auto h-[57vh] rounded-tl-lg rounded-tr-lg">
              <table className="table w-full">
                <thead>
                  <tr className="border-b border-slate-200">
                    {/* <th>{""}</th> */}
                    <th>Name</th>
                    <th>Email</th>
                    <th>Role</th>
                    <th>{""}</th>
                  </tr>
                </thead>
                <tbody>
                  {data[0].members ? (
                    <>
                      {data[0].members.map((member: any, index: number) => (
                        <tr key={index} className="hover:bg-slate-50">
                          {/* <th className="w-4">
                        <input type="checkbox" className="checkbox" />
                      </th> */}
                          <td className="lg:w-96 2xl:w-[56rem]">
                            <div className="flex items-center space-x-3">
                              <div className="mask w-12 flex items-center">
                                <AvatarIcon intent="org">
                                  {member.name.charAt(0).toUpperCase() || "U"}
                                </AvatarIcon>
                              </div>
                              <div>
                                <div className="text-slate-600">
                                  {member.name || "Name"}
                                </div>
                                <div className="text-sm text-slate-400">
                                  {member.handle || "username"}
                                </div>
                              </div>
                            </div>
                          </td>
                          <td className="lg:w-96 2xl:w-[56rem]">
                            {member.email || "abc@gmail.com"}
                          </td>
                          <td>
                            <Form
                              method="post"
                              reloadDocument
                              onChange={roleChange}
                            >
                              <input
                                className="hidden"
                                type="text"
                                name="orgid"
                                defaultValue={data[0].uuid}
                              />
                              <input
                                className="hidden"
                                type="text"
                                name="roleEmail"
                                defaultValue={member.email}
                              />
                              <SelectPrimitive.Root name="role">
                                <SelectPrimitive.Trigger
                                  asChild
                                  aria-label="Role"
                                  className="focus:outline-none"
                                >
                                  <button className="flex foxus:outline-none">
                                    <SelectPrimitive.Icon className="capitalize focus:outline-none flex justify-center items-center text-sm gap-x-2 text-slate-600 border border-slate-200 rounded px-2 py-1">
                                      <div className="hidden">{role}</div>
                                      {member.role}{" "}
                                      <ChevronDown className="w-4" />
                                    </SelectPrimitive.Icon>
                                  </button>
                                </SelectPrimitive.Trigger>
                                <SelectPrimitive.Content
                                  position="popper"
                                  sideOffset={7}
                                  align="start"
                                  className="z-50"
                                >
                                  <SelectPrimitive.Viewport className="bg-white rounded shadow-lg">
                                    <SelectPrimitive.Group>
                                      {allRoles.map(
                                        (role: any, index: number) => (
                                          <SelectPrimitive.Item
                                            key={`${role}-${index}`}
                                            value={role.value}
                                            className="flex items-center w-24 px-2 py-2 rounded text-sm text-slate-600 cursor-pointer hover:bg-slate-100 hover:border-none focus:outline-none"
                                          >
                                            <SelectPrimitive.ItemText className="text-sm">
                                              {role.label}
                                            </SelectPrimitive.ItemText>
                                          </SelectPrimitive.Item>
                                        )
                                      )}
                                    </SelectPrimitive.Group>
                                  </SelectPrimitive.Viewport>
                                </SelectPrimitive.Content>
                              </SelectPrimitive.Root>
                            </Form>
                          </td>
                          <th className="lg:w-96">
                            <Modal
                              btnName={
                                <button className="font-normal">Remove</button>
                              }
                              title="Remove member"
                            >
                              <Form
                                method="post"
                                reloadDocument
                                className="w-full"
                              >
                                <input
                                  className="hidden"
                                  type="text"
                                  name="orgid"
                                  defaultValue={data[0].uuid}
                                />
                                <input
                                  className="hidden"
                                  type="text"
                                  name="removeEmail"
                                  defaultValue={member.email}
                                />
                                <div className="text-slate-600">
                                  Are you sure you want to remove this user?
                                </div>
                                <div className="pt-12 grid justify-items-end w-full">
                                  <div className="flex justify-end w-1/2">
                                    <Button
                                      intent="primary"
                                      type="submit"
                                      className="pl-2"
                                      fullWidth={false}
                                    >
                                      Remove
                                    </Button>
                                  </div>
                                </div>
                              </Form>
                            </Modal>
                          </th>
                        </tr>
                      ))}
                    </>
                  ) : (
                    <div>No members found</div>
                  )}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </Suspense>
  );
}

// ############################ error boundary ###########################

export function ErrorBoundary() {
  return (
    <div className="flex flex-col h-screen justify-center items-center bg-slate-50">
      <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
      <div className="text-3xl text-slate-600 font-medium">
        Something went wrong :(
      </div>
      <img src="/error/ErrorFunction.gif" alt="Error" width="500" />
    </div>
  );
}

export function CatchBoundary() {
  return (
    <div className="flex flex-col h-screen justify-center items-center bg-slate-50">
      <div className="text-3xl text-slate-600 font-medium">Oops!!</div>
      <div className="text-3xl text-slate-600 font-medium">
        Something went wrong :(
      </div>
      <img src="/error/ErrorFunction.gif" alt="Error" width="500" />
    </div>
  );
}
