import styles from "./style.module.css";

export default function Table({ rows }: { rows: Array<Array<string>> }) {
  return (
    <div
      className={
        "-mx-6 mb-4 mt-6 overflow-x-auto overscroll-x-contain px-6 pb-4 " +
        styles.container
      }
    >
      <table className="w-full border-collapse text-sm">
        <tbody className="align-baseline text-gray-900 dark:text-gray-100">
          {rows.map((row, i) => (
            <tr
              key={`${i}`}
              className="border-b border-gray-100 dark:border-neutral-700/50 pr-4"
            >
              {row.map((val, k) => (
                <td
                  key={`${i}_${k}`}
                  className={`py-2 pl-4 ${k === 0 ? "w-32 font-semibold" : ""}`}
                >
                  {val}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
