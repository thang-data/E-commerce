import { FormContext } from "vee-validate";

export default function useFormHandler<
  T extends Record<string, string>,
  F extends (...args: any) => Promise<any>
>(
  form: FormContext<T> | any,
  onSubmit: F
): {
  formValues: T;
  formHandle: () => ReturnType<F>;
} {
  const { values, meta, submitForm } = form;

  const formHandle = async function () {
    await submitForm();

    if (!meta.value.valid) return;

    const formData = new FormData();
    for (const key in values) {
      formData.append(key, values[key]);
    }
    return await onSubmit(formData);
  } as any;

  return {
    formValues: values,
    formHandle,
  };
}
