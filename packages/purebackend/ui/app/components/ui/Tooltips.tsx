import * as Tooltip from "@radix-ui/react-tooltip";

function Tooltips({ msg, children }: any) {
  return (
    <Tooltip.Provider>
      <Tooltip.Root>
        <Tooltip.Trigger>{children}</Tooltip.Trigger>
        <Tooltip.Portal>
          <Tooltip.Content
            className="bg-slate-200 text-slate-600 p-1 rounded-lg"
            sideOffset={25}
            side="right"
          >
            {msg}
            <Tooltip.Arrow className="fill-slate-200" />
          </Tooltip.Content>
        </Tooltip.Portal>
      </Tooltip.Root>
    </Tooltip.Provider>
  );
}

export default Tooltips;
