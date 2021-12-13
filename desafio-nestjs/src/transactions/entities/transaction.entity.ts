import { Column, DataType, Model, PrimaryKey, Table } from "sequelize-typescript";

@Table({
    tableName:'transactions',
    createdAt: 'created_at',
    updatedAt: 'updated_at'
})
export class Transaction extends Model {
    @PrimaryKey
    @Column({ type: DataType.UUID, defaultValue: DataType.UUIDV4 })
    id: string;

    @Column({ allowNull: false })
    account_id: string;

    @Column({ allowNull: false, type: DataType.DECIMAL(10, 2) })
    amount: number;
}
