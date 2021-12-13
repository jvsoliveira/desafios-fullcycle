import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import { CreateTransactionDto } from './dto/create-transaction.dto';
import { UpdateTransactionDto } from './dto/update-transaction.dto';
import { Transaction } from './entities/transaction.entity';

@Injectable()
export class TransactionsService {

constructor(
  @InjectModel(Transaction)
  private transactionModel: typeof Transaction
  ){}


  create(createTransactionDto: CreateTransactionDto) {
    return this.transactionModel.create(createTransactionDto);
  }

  findAll() {
    return this.transactionModel.findAll();
  }

  findOne(id: string) {
    return this.transactionModel.findByPk(id);  }

  async update(id: string, updateTransactionDto: UpdateTransactionDto) {
    const transaction = await this.findOne(id);
    return transaction.update(updateTransactionDto)
  }

  async remove(id: string) {
    const transaction = await this.findOne(id);
    return transaction.destroy();
  }
}
