export interface ProfileFieldProps {
  label: string
  value: string
}

export default function ProfileField({ label, value }: ProfileFieldProps) {
  return (
    <div className="pl-profile-field">
      <dt className="pl-profile-field__label">{label}</dt>
      <dd className="pl-profile-field__value">{value}</dd>
    </div>
  )
}
