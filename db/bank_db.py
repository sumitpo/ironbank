from typing import List, Optional

from sqlalchemy import Float, ForeignKeyConstraint, Index, Text
from sqlalchemy.dialects.mysql import INTEGER
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column, relationship

class Base(DeclarativeBase):
    pass


class District(Base):
    __tablename__ = 'district'

    district_id: Mapped[int] = mapped_column(INTEGER(11), primary_key=True)
    district_name: Mapped[Optional[str]] = mapped_column(Text)
    region: Mapped[Optional[str]] = mapped_column(Text)
    A4: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    A5: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    A6: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    A7: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    A8: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    A9: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    A10: Mapped[Optional[float]] = mapped_column(Float)
    A11: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    A12: Mapped[Optional[float]] = mapped_column(Float)
    A13: Mapped[Optional[float]] = mapped_column(Float)
    A14: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    A15: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    A16: Mapped[Optional[int]] = mapped_column(INTEGER(11))

    account: Mapped[List['Account']] = relationship('Account', back_populates='district')
    client: Mapped[List['Client']] = relationship('Client', back_populates='district')


class Account(Base):
    __tablename__ = 'account'
    __table_args__ = (
        ForeignKeyConstraint(['district_id'], ['district.district_id'], name='account_ibfk_1'),
        Index('account_ibfk_1', 'district_id'),
        Index('account_id', 'account_id', unique=True)
    )

    account_id: Mapped[int] = mapped_column(INTEGER(11), primary_key=True)
    district_id: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    frequency: Mapped[Optional[str]] = mapped_column(Text)
    date: Mapped[Optional[int]] = mapped_column(INTEGER(11))

    district: Mapped[Optional['District']] = relationship('District', back_populates='account')
    disp: Mapped[List['Disp']] = relationship('Disp', back_populates='account')
    loan: Mapped[List['Loan']] = relationship('Loan', back_populates='account')
    order: Mapped[List['Order']] = relationship('Order', back_populates='account')
    trans: Mapped[List['Trans']] = relationship('Trans', back_populates='account_')


class Client(Base):
    __tablename__ = 'client'
    __table_args__ = (
        ForeignKeyConstraint(['district_id'], ['district.district_id'], name='client_ibfk_1'),
        Index('client_ibfk_1', 'district_id'),
        Index('client_id', 'client_id', unique=True)
    )

    client_id: Mapped[int] = mapped_column(INTEGER(11), primary_key=True)
    district_id: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    birth_number: Mapped[Optional[int]] = mapped_column(INTEGER(11))

    district: Mapped[Optional['District']] = relationship('District', back_populates='client')
    disp: Mapped[List['Disp']] = relationship('Disp', back_populates='client')


class Disp(Base):
    __tablename__ = 'disp'
    __table_args__ = (
        ForeignKeyConstraint(['account_id'], ['account.account_id'], name='disp_ibfk_2'),
        ForeignKeyConstraint(['client_id'], ['client.client_id'], name='disp_ibfk_1'),
        Index('disp_ibfk_1', 'client_id'),
        Index('disp_ibfk_2', 'account_id'),
        Index('disp_id', 'disp_id', unique=True)
    )

    disp_id: Mapped[int] = mapped_column(INTEGER(11), primary_key=True)
    client_id: Mapped[int] = mapped_column(INTEGER(11))
    account_id: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    type: Mapped[Optional[str]] = mapped_column(Text)

    account: Mapped[Optional['Account']] = relationship('Account', back_populates='disp')
    client: Mapped['Client'] = relationship('Client', back_populates='disp')
    card: Mapped[List['Card']] = relationship('Card', back_populates='disp')


class Loan(Base):
    __tablename__ = 'loan'
    __table_args__ = (
        ForeignKeyConstraint(['account_id'], ['account.account_id'], name='loan_ibfk_1'),
        Index('loan_ibfk_1', 'account_id')
    )

    loan_id: Mapped[int] = mapped_column(INTEGER(11), primary_key=True)
    account_id: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    date: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    amount: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    duration: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    payments: Mapped[Optional[float]] = mapped_column(Float)
    status: Mapped[Optional[str]] = mapped_column(Text)

    account: Mapped[Optional['Account']] = relationship('Account', back_populates='loan')


class Order(Base):
    __tablename__ = 'order'
    __table_args__ = (
        ForeignKeyConstraint(['account_id'], ['account.account_id'], name='order_ibfk_1'),
        Index('order_ibfk_1', 'account_id')
    )

    order_id: Mapped[int] = mapped_column(INTEGER(11), primary_key=True)
    account_id: Mapped[int] = mapped_column(INTEGER(11))
    bank_to: Mapped[Optional[str]] = mapped_column(Text)
    account_to: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    amount: Mapped[Optional[float]] = mapped_column(Float)
    k_symbol: Mapped[Optional[str]] = mapped_column(Text)

    account: Mapped['Account'] = relationship('Account', back_populates='order')


class Trans(Base):
    __tablename__ = 'trans'
    __table_args__ = (
        ForeignKeyConstraint(['account_id'], ['account.account_id'], name='trans_ibfk_1'),
        Index('trans_ibfk_1', 'account_id')
    )

    trans_id: Mapped[int] = mapped_column(INTEGER(11), primary_key=True)
    account_id: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    date: Mapped[Optional[int]] = mapped_column(INTEGER(11))
    type: Mapped[Optional[str]] = mapped_column(Text)
    operation: Mapped[Optional[str]] = mapped_column(Text)
    amount: Mapped[Optional[float]] = mapped_column(Float)
    balance: Mapped[Optional[float]] = mapped_column(Float)
    k_symbol: Mapped[Optional[str]] = mapped_column(Text)
    bank: Mapped[Optional[str]] = mapped_column(Text)
    account: Mapped[Optional[int]] = mapped_column(INTEGER(11))

    account_: Mapped[Optional['Account']] = relationship('Account', back_populates='trans')


class Card(Base):
    __tablename__ = 'card'
    __table_args__ = (
        ForeignKeyConstraint(['disp_id'], ['disp.disp_id'], name='card_ibfk_1'),
        Index('card_ibfk_1', 'disp_id'),
        Index('card_id', 'card_id', unique=True)
    )

    card_id: Mapped[int] = mapped_column(INTEGER(11), primary_key=True)
    disp_id: Mapped[int] = mapped_column(INTEGER(11))
    type: Mapped[Optional[str]] = mapped_column(Text)
    issued: Mapped[Optional[str]] = mapped_column(Text)

    disp: Mapped['Disp'] = relationship('Disp', back_populates='card')
