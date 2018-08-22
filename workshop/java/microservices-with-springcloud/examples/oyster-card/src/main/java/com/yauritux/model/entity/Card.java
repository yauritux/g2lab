package com.yauritux.model.entity;

import java.math.BigDecimal;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

import com.yauritux.model.builder.CardBuilder;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@Entity
@Table(name = "cards")
public class Card {

	@Id
	@GeneratedValue(strategy = GenerationType.AUTO)
	private Long id;
	
	@Column(name = "serial_no", nullable = false, unique = true)
	private String serialNo;
	
	@Column(name = "owner", nullable = true)
	private String owner;
	
	@Column(name = "balance")
	private BigDecimal balance;
	
	public Card() {}
	
	public Card(CardBuilder cardBuilder) {
		this.owner = cardBuilder.getOwner();
		this.serialNo = cardBuilder.getSerialNo();
		this.balance = cardBuilder.getBalance();
	}
	
	public Long getId() {
		return id;
	}
	
	protected void setId(Long id) {
		this.id = id;
	}
	
	public String getSerialNo() {
		return serialNo;
	}
	
	public void setSerialNo(String serialNo) {
		this.serialNo = serialNo;
	}
	
	public String getOwner() {
		return owner;
	}
	
	public void setOwner(String owner) { 
		this.owner = owner;
	}
	
	public BigDecimal getBalance() {
		return balance;
	}
	
	public void setBalance(BigDecimal balance) {
		this.balance = balance;
	}
}
