package com.yauritux.model.builder;

import java.math.BigDecimal;
import java.util.UUID;

import com.yauritux.model.entity.Card;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public class CardBuilder {

	private final String owner;
	private final String serialNo;
	private BigDecimal balance;
	
	public CardBuilder(String owner) {
		this.owner = owner;
		this.serialNo = UUID.randomUUID().toString();
	}
	
	public final String getOwner() {
		return owner;
	}
	
	public final String getSerialNo() {
		return serialNo;
	}
	
	public BigDecimal getBalance() {
		return balance;
	}
	
	public CardBuilder setBalance(BigDecimal balance) {
		this.balance = balance;
		return this;
	}
	
	public Card build() {
		return new Card(this);
	}
}
