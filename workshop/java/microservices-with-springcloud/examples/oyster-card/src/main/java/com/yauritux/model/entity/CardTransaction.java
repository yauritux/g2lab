package com.yauritux.model.entity;

import java.util.Date;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.FetchType;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;
import javax.persistence.Temporal;
import javax.persistence.TemporalType;

import com.yauritux.model.constant.BarrierType;
import com.yauritux.model.constant.JourneyType;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@Entity
@Table(name = "card_transactions")
public class CardTransaction {

	@Id
	@GeneratedValue(strategy = GenerationType.AUTO)	
	private Long id;
	
	@Column(name = "trx_date")
	@Temporal(TemporalType.TIMESTAMP)
	private Date date;
	
	@Column(name = "journey_type")
	@Enumerated(EnumType.STRING)
	private JourneyType journeyType;
	
	@Column(name = "barrier_type", nullable = true)
	@Enumerated(EnumType.STRING)
	private BarrierType barrierType;
	
	@ManyToOne(fetch = FetchType.LAZY)
	@JoinColumn(name = "card_id", referencedColumnName = "id")
	private Card card;
	
	public CardTransaction() {}
	
	public Long getId() {
		return id;
	}
	
	protected void setId(Long id) {
		this.id = id;
	}
	
	public Date getDate() {
		return date;
	}
	
	public void setDate(Date date) {
		this.date = date;
	}
	
	public JourneyType getJourneyType() {
		return journeyType;
	}
	
	public void setJourneyType(JourneyType journeyType) {
		this.journeyType = journeyType;
	}
	
	public BarrierType getBarrierType() {
		return barrierType;
	}
	
	public void setBarrierType(BarrierType barrierType) { 
		this.barrierType = barrierType;
	}
	
	public Card getCard() {
		return card;
	}
	
	public void setCard(Card card) {
		this.card = card;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + ((id == null) ? 0 : id.hashCode());
		return result;
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		CardTransaction other = (CardTransaction) obj;
		if (id == null) {
			if (other.id != null)
				return false;
		} else if (!id.equals(other.id))
			return false;
		return true;
	}
}
