import { useLoaderData } from "@remix-run/react";
import {
  fetchPublicProfile,
  fetchUserSettings,
} from "~/routes/api/auth.server";
import Avatar from "~/components/ui/Avatar";
import ProfileCard from "~/components/ProfileCard";
import NavBar from "~/components/Navbar";
import type { MetaFunction } from "@remix-run/node";
import Error from "~/Error404";
import { getSession } from "~/session";
import { Suspense } from "react";
import Loader from "~/components/ui/Loading";
import { fetchOrgDetails } from "./api/org.server";
import CalendarHeatmap from "react-calendar-heatmap";

export const meta: MetaFunction = () => ({
  charset: "utf-8",
  title: "Profile | PureML",
  viewport: "width=device-width,initial-scale=1",
});

export async function loader({ params, request }: any) {
  const session = await getSession(request.headers.get("Cookie"));
  const publicProfile = await fetchPublicProfile(params.username);
  const accesstoken = session.get("accessToken");
  if (accesstoken) {
    const userProfile = await fetchUserSettings(accesstoken);
    const orgId = session.get("orgId");
    const orgDetails = await fetchOrgDetails(orgId, session.get("accessToken"));
    return { userProfile, publicProfile, orgDetails };
  }
  if (!publicProfile) return null;
  return publicProfile;
}

const today = new Date();

function shiftDate(date: Date, numDays: number) {
  const newDate = new Date(date);
  newDate.setDate(newDate.getDate() + numDays);
  return newDate;
}

function getRange(count: number) {
  return Array.from({ length: count }, (_, i) => i);
}

function getRandomInt(min: number, max: number) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

export default function UserProfile() {
  const userProfileData = useLoaderData();
  const randomValues = getRange(270).map((index: any) => {
    return {
      date: shiftDate(today, -index),
      count: getRandomInt(1, 3),
    };
  });

  // ##### public profile for logged in users ####

  if (userProfileData.userProfile && userProfileData.publicProfile) {
    if (userProfileData.publicProfile[0]) {
      return (
        <Suspense fallback={<Loader />}>
          <div className="flex flex-col w-screen h-screen overflow-hidden">
            {userProfileData.userProfile ? (
              <NavBar
                intent="loggedIn"
                user={userProfileData.userProfile[0].name
                  .charAt(0)
                  .toUpperCase()}
                orgName={
                  <a
                    href={`/org/${userProfileData.orgDetails[0].name}`}
                    className="flex items-center justify-center"
                  >
                    {userProfileData.orgDetails[0].name}
                  </a>
                }
                orgAvatarName={userProfileData.orgDetails[0].name.charAt(0)}
              />
            ) : (
              <NavBar
                intent="loggedOut"
                user=""
                orgName={
                  <a
                    href="/models"
                    className="flex items-center justify-center pr-8"
                  >
                    <img
                      src="/PureMLLogoWText.svg"
                      alt="Logo"
                      className="w-20"
                    />
                  </a>
                }
                orgAvatarName=""
              />
            )}

            <div className="flex justify-center w-full">
              <div className="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl">
                <div className="flex flex-col gap-y-8 md:flex md:flex-row px-12 pt-8 pb-12 text-slate-600 font-medium overflow-auto">
                  <div className="h-full md:w-28 md:w-44 lg:w-56 2xl:w-96">
                    <div className="h-28 w-28 md:h-36 md:w-36 lg:w-56 lg:h-56 2xl:h-96 2xl:w-96 flex items-center justify-center text-md text-blue-600 bg-blue-200 rounded-lg">
                      <Avatar intent="profile">
                        {userProfileData.publicProfile[0]?.name
                          .charAt(0)
                          .toUpperCase() || "User"}
                      </Avatar>
                    </div>
                    <div className="pt-6 font-medium text-base text-slate-600">
                      {userProfileData.publicProfile[0]?.name || "Name"}
                    </div>
                    <div className="pb-6 text-base font-normal">
                      {userProfileData.publicProfile[0]?.email || "Email"}
                    </div>
                    {/* <Button aria-label="follow" intent="primary" >
                Follow
              </Button> */}
                    <div className="text-base">
                      <span>Bio</span>
                    </div>
                    <div className="font-normal text-base text-slate-600">
                      {userProfileData.publicProfile[0]?.bio || "Add your bio"}
                    </div>
                    {/* <div className="text-base pt-8">Organizations</div>
            {userProfileData.publicProfile[0].orgs.length !== 0 ? (
              <div>
                {userProfileData.publicProfile[0].orgs.map((org: any) => (
                  <Button intent="org"  key={org.name}>
                    {org.name}
                  </Button>
                ))}
              </div>
            ) : (
              "-"
            )} */}
                  </div>
                  <div className="md:pl-12 w-full md:w-3/4 2xl:w-1/2">
                    {/* <div>
                  <div className="flex justify-between">
                    <div className="text-slate-800 font-medium font-base">
                      Total contributions this year
                    </div>
                  </div>
                  <div className="pt-6">
                    <CalendarHeatmap
                      startDate={shiftDate(today, -270)}
                      endDate={today}
                      values={randomValues}
                      classForValue={(value: any) => {
                        if (!value) {
                          return "color-empty";
                        }
                        return `color-github-${value.count}`;
                      }}
                      showWeekdayLabels={true}
                    />
                  </div>
                  <div className="flex justify-end"><img src="/imgs/ContributionScale.png" /></div>
                </div> */}
                    <div className="font-medium">Overview</div>
                    <div className="pt-8 flex flex-col gap-y-4 md:flex md:flex-row w-full md:gap-x-4">
                      <div className="w-full">
                        <ProfileCard
                          title="Models"
                          count={
                            userProfileData.publicProfile[0]
                              ?.number_of_models || "0"
                          }
                        />
                      </div>
                      <div className="w-full">
                        <ProfileCard
                          title="Datasets"
                          count={
                            userProfileData.publicProfile[0]
                              ?.number_of_datasets || "0"
                          }
                        />
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </Suspense>
      );
    }
  }

  // ##### anonymous view of public profile ####

  if (userProfileData[0]) {
    return (
      <Suspense fallback={<Loader />}>
        <div className="flex flex-col w-screen h-screen overflow-hidden">
          <NavBar
            intent="loggedOut"
            user=""
            orgName={
              <a href="/models" className="flex items-center justify-center">
                <img src="/PureMLLogoWText.svg" alt="Logo" className="w-20" />
              </a>
            }
            orgAvatarName=""
          />
          <div className="flex justify-center w-full">
            <div className="flex flex-col max-w-screen-2xl gap-y-8 md:flex md:flex-row px-12 pt-8 pb-12 text-slate-600 font-medium overflow-auto">
              <div className="h-full md:w-28 md:w-44 lg:w-56 2xl:w-96">
                <div className="h-28 w-28 md:h-36 md:w-36 lg:w-56 lg:h-56 2xl:h-96 2xl:w-96 flex items-center justify-center text-md text-blue-600 bg-blue-200 rounded-lg">
                  <Avatar intent="profile">
                    {userProfileData[0]?.name.charAt(0).toUpperCase() || "User"}
                  </Avatar>
                </div>
                <div className="pt-6 font-medium text-base text-slate-600">
                  {userProfileData[0]?.name || "Name"}
                </div>
                <div className="pb-6 text-base font-normal">
                  {userProfileData[0]?.email || "Email"}
                </div>
                {/* <Button aria-label="follow" intent="primary" >
              Follow
            </Button> */}
                <div className="text-base">
                  <span>Bio</span>
                </div>
                <div className="font-medium text-base text-slate-600">
                  {userProfileData[0]?.bio || "Add your bio"}
                </div>
                {/* <div className="text-base pt-8">Organizations</div>
            {userProfileData[0].orgs.length !== 0 ? (
              <div>
                {userProfileData[0].orgs.map((org: any) => (
                  <Button intent="org"  key={org.name}>
                    {org.name}
                  </Button>
                ))}
              </div>
            ) : (
              "-"
            )} */}
              </div>
              <div className="md:pl-12 w-full md:w-3/4 2xl:w-1/2">
                {/* <div>
                <div className="flex justify-between">
                  <div className="text-slate-800 font-medium font-base">
                    Total contributions this year
                  </div>
                </div>
                <div className="pt-6">
                  <CalendarHeatmap
                    startDate={shiftDate(today, -270)}
                    endDate={today}
                    values={randomValues}
                    classForValue={(value: any) => {
                      if (!value) {
                        return "color-empty";
                      }
                      return `color-github-${value.count}`;
                    }}
                    showWeekdayLabels={true}
                  />
                </div>
                <div className="flex justify-end"><img src="/imgs/ContributionScale.png" /></div>
              </div> */}
                <div className="font-medium">Overview</div>
                <div className="pt-8 flex flex-col gap-y-4 md:flex md:flex-row w-full md:gap-x-4">
                  <div className="w-full">
                    <ProfileCard
                      title="Models"
                      count={userProfileData[0]?.number_of_models || "0"}
                    />
                  </div>
                  <div className="w-full">
                    <ProfileCard
                      title="Datasets"
                      count={userProfileData[0]?.number_of_datasets || "0"}
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </Suspense>
    );
  }

  // ##### when profile does not exist ####

  return (
    <Suspense fallback={<Loader />}>
      {userProfileData.userProfile ? (
        <NavBar
          intent="loggedIn"
          user={userProfileData.userProfile[0].name.charAt(0).toUpperCase()}
          orgName={
            <a
              href={`/org/${userProfileData.orgDetails[0].name}`}
              className="flex items-center justify-center"
            >
              {userProfileData.orgDetails[0].name}
            </a>
          }
          orgAvatarName={userProfileData.orgDetails[0].name.charAt(0)}
        />
      ) : (
        <NavBar
          intent="loggedOut"
          user=""
          orgName={
            <a href="/models" className="flex items-center justify-center pr-8">
              <img src="/PureMLLogoWText.svg" alt="Logo" className="w-20" />
            </a>
          }
          orgAvatarName=""
        />
      )}
      <Error />
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
